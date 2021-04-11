package zkaffg

import (
	"math/big"
	"testing"

	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/sample"
	"github.com/taurusgroup/cmp-ecdsa/pkg/paillier"
	"github.com/taurusgroup/cmp-ecdsa/wip/zkcommon"
)

func TestAffG(t *testing.T) {
	verifierPaillier := zkcommon.VerifierPaillierPublic
	verifierPedersen := zkcommon.Pedersen
	prover := zkcommon.ProverPaillierPublic
	x := sample.IntervalL()
	y := sample.IntervalLPrime()
	c := big.NewInt(12)

	C, _ := verifierPaillier.Enc(c, nil)
	var X curve.Point
	X.ScalarBaseMult(curve.NewScalarBigInt(x))
	Y, rhoY := prover.Enc(y, nil)

	var tmp paillier.Ciphertext
	tmp.Mul(verifierPaillier, C, x)
	D, rho := verifierPaillier.Enc(y, nil)
	D.Add(verifierPaillier, D, &tmp)

	proof := NewProof(prover, verifierPaillier, verifierPedersen, D, C, Y, &X, x, y, rhoY, rho)
	if !proof.Verify(prover, verifierPaillier, verifierPedersen, D, C, Y, &X) {
		t.Error("failed to verify")
	}
}
