package zkaffg

import (
	"math/big"

	"github.com/taurusgroup/cmp-ecdsa/pb"
	"github.com/taurusgroup/cmp-ecdsa/pkg/hash"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/arith"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/sample"
	"github.com/taurusgroup/cmp-ecdsa/pkg/paillier"
	"github.com/taurusgroup/cmp-ecdsa/pkg/pedersen"
	"github.com/taurusgroup/cmp-ecdsa/pkg/zk"
)

type (
	Public struct {
		// C = Enc₀(?;?)
		// Kⱼ = Encⱼ(kⱼ; )
		C *paillier.Ciphertext

		// D = (x ⨀ C) ⨁ Enc₀(y;ρ)
		//   = (xᵢ ⨀ Kⱼ) ⨁ Encⱼ(- βᵢⱼ; sᵢⱼ)
		D *paillier.Ciphertext

		// Y = Enc₁(y;ρ')
		//   = Encᵢ(βᵢⱼ,rᵢⱼ)
		// y is Bob's additive share
		Y *paillier.Ciphertext

		// X = gˣ
		// x is Alice's multiplicative share
		X *curve.Point

		// Prover = N₁
		// Verifier = N₀
		Prover, Verifier *paillier.PublicKey
		Aux              *pedersen.Parameters
	}
	Private struct {
		// X ∈ ± 2ˡ
		// Bob's multiplicative share
		X *big.Int

		// Y ∈ ± 2ˡº
		// Bob's additive share βᵢⱼ
		Y *big.Int

		// Rho = Nonce D = sᵢⱼ
		Rho *big.Int

		// RhoY = Nonce Y = rᵢⱼ
		RhoY *big.Int
	}
)

type Commitment struct {
	// A = (alpha ⊙ c ) ⊕ Enc(N0, beta, r)
	A *paillier.Ciphertext

	// Bx = Enc(N1, alpha, rx)
	Bx *curve.Point

	// By = Enc(N1, beta, ry)
	By *paillier.Ciphertext

	// E = s^alpha t^gamma
	// S = sˣ t^m
	E, S *big.Int

	// F = s^beta t^delta
	// T = sʸ t^mu
	F, T *big.Int
}

func (public Public) Prove(hash *hash.Hash, private Private) (*pb.ZKAffG, error) {
	N0 := public.Verifier.N
	N1 := public.Prover.N

	verifier := public.Verifier
	prover := public.Prover

	alpha := sample.IntervalLEps()
	beta := sample.IntervalLPrimeEps()

	r := sample.UnitModN(N0)
	rY := sample.UnitModN(N1)

	gamma := sample.IntervalLEpsN()
	m := sample.IntervalLEpsN()
	delta := sample.IntervalLEpsN()
	mu := sample.IntervalLN()

	A, _ := verifier.Enc(beta, r)
	tmpCt := paillier.NewCiphertext().Mul(verifier, public.C, alpha)
	A.Add(verifier, A, tmpCt)

	var Bx curve.Point
	Bx.ScalarBaseMult(curve.NewScalarBigInt(alpha))
	By, _ := prover.Enc(beta, rY)

	commitment := Commitment{
		A:  A,
		Bx: &Bx,
		By: By,
		E:  public.Aux.Commit(alpha, gamma),
		S:  public.Aux.Commit(private.X, m),
		F:  public.Aux.Commit(beta, delta),
		T:  public.Aux.Commit(private.Y, mu),
	}

	e, err := challenge(hash, public, commitment)
	if err != nil {
		return nil, err
	}

	return &pb.ZKAffG{
		A:  pb.NewCiphertext(A),
		Bx: pb.NewPoint(&Bx),
		By: pb.NewCiphertext(By),
		E:  pb.NewInt(commitment.E),
		S:  pb.NewInt(commitment.S),
		F:  pb.NewInt(commitment.F),
		T:  pb.NewInt(commitment.T),
		Z1: zk.Affine(alpha, e, private.X),
		Z2: zk.Affine(beta, e, private.Y),
		Z3: zk.Affine(gamma, e, m),
		Z4: zk.Affine(delta, e, mu),
		W:  zk.AffineNonce(r, private.Rho, e, verifier),
		Wy: zk.AffineNonce(rY, private.RhoY, e, prover),
	}, nil
}
func (public Public) Verify(hash *hash.Hash, proof *pb.ZKAffG) bool {
	if !proof.IsValid() {
		return false
	}

	verifier := public.Verifier
	prover := public.Prover

	z1, z2, z3, z4 := proof.Z1.Unmarshal(), proof.Z2.Unmarshal(), proof.Z3.Unmarshal(), proof.Z4.Unmarshal()
	w, wY := proof.GetW().Unmarshal(), proof.GetWy().Unmarshal()

	if !arith.IsInIntervalLEps(z1) {
		return false
	}
	if !arith.IsInIntervalLPrimeEps(z2) {
		return false
	}

	Bx, err := proof.GetBx().Unmarshal()
	if err != nil {
		return false
	}

	A, By := proof.GetA().Unmarshal(), proof.By.Unmarshal()

	E, S, F, T := proof.GetE().Unmarshal(), proof.GetS().Unmarshal(), proof.GetF().Unmarshal(), proof.GetT().Unmarshal()

	e, err := challenge(hash, public, Commitment{
		A:  A,
		Bx: Bx,
		By: By,
		E:  E,
		S:  S,
		F:  F,
		T:  T,
	})
	if err != nil {
		return false
	}

	lhsCt, rhsCt := paillier.NewCiphertext(), paillier.NewCiphertext()
	{
		// lhsCt = z1•c + Enc_N0(z2;w)
		rhsCt.Enc(verifier, z2, w)
		lhsCt.Mul(verifier, public.C, z1)
		lhsCt.Add(verifier, lhsCt, rhsCt)

		// rhsCt = A + e•D
		rhsCt.Mul(verifier, public.D, e)
		rhsCt.Add(verifier, rhsCt, A)

		if !lhsCt.Equal(rhsCt) {
			return false
		}
	}

	{
		var lhsPt, rhsPt curve.Point
		// lhsPt = [z₁]G
		lhsPt.ScalarBaseMult(curve.NewScalarBigInt(z1))

		// rhsPt = Bₓ + [e]G
		rhsPt.ScalarBaseMult(curve.NewScalarBigInt(e))
		rhsPt.Add(&rhsPt, Bx)
		if lhsPt.Equal(&rhsPt) {
			return false
		}
	}

	{
		// lhsCt = Enc₁(z₂; wy)
		lhsCt.Enc(prover, z2, wY)

		// rhsCt = By ⊕ (e ⊙ Y)
		rhsCt.Mul(prover, public.Y, e)
		rhsCt.Add(prover, rhsCt, By)

		if !lhsCt.Equal(rhsCt) {
			return false
		}
	}

	if !public.Aux.Verify(z1, z3, E, S, e) {
		return false
	}

	if !public.Aux.Verify(z2, z4, F, T, e) {
		return false
	}

	return true
}

func challenge(hash *hash.Hash, public Public, commitment Commitment) (*big.Int, error) {

	err := hash.WriteAny(public.Aux, public.Prover, public.Verifier,
		public.C, public.D, public.Y, public.X,
		commitment.A, commitment.Bx, commitment.By,
		commitment.E, commitment.S, commitment.F, commitment.T)
	if err != nil {
		return nil, err
	}

	return hash.ReadFqNegative()
}
