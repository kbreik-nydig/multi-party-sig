package zkaffg

import (
	"math/big"
	"testing"

	"github.com/taurusgroup/cmp-ecdsa/pb"
	"github.com/taurusgroup/cmp-ecdsa/pkg/hash"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/sample"
	"github.com/taurusgroup/cmp-ecdsa/pkg/paillier"
	"github.com/taurusgroup/cmp-ecdsa/pkg/zk"
	"google.golang.org/protobuf/proto"
)

func TestAffG(t *testing.T) {
	verifierPaillier := zk.VerifierPaillierPublic
	verifierPedersen := zk.Pedersen
	prover := zk.ProverPaillierPublic

	c := big.NewInt(12)
	C, _ := verifierPaillier.Enc(c, nil)

	var X curve.Point
	x := sample.IntervalL()
	X.ScalarBaseMult(curve.NewScalarBigInt(x))

	y := sample.IntervalLPrime()
	Y, rhoY := prover.Enc(y, nil)

	tmp := paillier.NewCiphertext()
	tmp.Mul(verifierPaillier, C, x)
	D, rho := verifierPaillier.Enc(y, nil)
	D.Add(verifierPaillier, D, tmp)

	public := Public{
		C:        C,
		D:        D,
		Y:        Y,
		X:        &X,
		Prover:   prover,
		Verifier: verifierPaillier,
		Aux:      verifierPedersen,
	}
	private := Private{
		X:    x,
		Y:    y,
		Rho:  rho,
		RhoY: rhoY,
	}
	proof, err := public.Prove(hash.New(nil), private)
	if err != nil {
		t.Error(err)
		return
	}
	out, err := proto.Marshal(proof)
	if err != nil {
		t.Error(err)
		return
	}
	proof2 := &pb.ZKAffG{}
	err = proto.Unmarshal(out, proof2)
	if err != nil {
		t.Error(err)
		return
	}

	if !public.Verify(hash.New(nil), proof2) {
		t.Error("failed to verify")
	}
}
