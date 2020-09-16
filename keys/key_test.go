package keys

import (
	"testing"

	"github.com/goat-systems/go-tezos/v3/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_NewKey(t *testing.T) {
	type want struct {
		wantErr     bool
		containsErr string
		secretKey   string
		publicKey   string
		address     string
	}

	cases := []struct {
		name  string
		input NewKeyInput
		want  want
	}{
		{
			"is successful with edesk",
			NewKeyInput{
				Esk:      "edesk1fddn27MaLcQVEdZpAYiyGQNm6UjtWiBfNP2ZenTy3CFsoSVJgeHM9pP9cvLJ2r5Xp2quQ5mYexW1LRKee2",
				Password: "password12345##",
				Kind:     Ed25519,
			},
			want{
				false,
				"",
				"edskRsPBsKuULoLTEQV2R9UbvSZbzFqvoESvp1mYyQJU8xi9mJamt88r5uTXbWQpVHjSiPWWtnoyqTCuSLQLxbEKUXfwwTccsF",
				"edpkuHMDkMz46HdRXYwom3xRwqk3zQ5ihWX4j8dwo2R2h8o4gPcbN5",
				"tz1L8fUQLuwRuywTZUP5JUw9LL3kJa8LMfoo",
			},
		},
		{
			"is successful with mnemonic",
			NewKeyInput{
				Kind:     Ed25519,
				Mnemonic: "normal dash crumble neutral reflect parrot know stairs culture fault check whale flock dog scout",
				Password: "PYh8nXDQLB",
				Email:    "vksbjweo.qsrgfvbw@tezos.example.org",
			},
			want{
				false,
				"",
				"edskRxB2DmoyZSyvhsqaJmw5CK6zYT7dbkUfEVSiQeWU1gw3ZMnC99QMMXru3imsbUrLhvuHktrymvNqhMxkhz7Y4LJAtevW5V",
				"edpkvEoAbkdaGALxi2FfeefB8hUkMZ4J1UVwkzyumx2GvbVpkYUHnm",
				"tz1Qny7jVMGiwRrP9FikRK95jTNbJcffTpx1",
			},
		},
		{
			"is successful with base58",
			NewKeyInput{
				Kind:          Ed25519,
				EncodedString: "edskRxB2DmoyZSyvhsqaJmw5CK6zYT7dbkUfEVSiQeWU1gw3ZMnC99QMMXru3imsbUrLhvuHktrymvNqhMxkhz7Y4LJAtevW5V",
			},
			want{
				false,
				"",
				"edskRxB2DmoyZSyvhsqaJmw5CK6zYT7dbkUfEVSiQeWU1gw3ZMnC99QMMXru3imsbUrLhvuHktrymvNqhMxkhz7Y4LJAtevW5V",
				"edpkvEoAbkdaGALxi2FfeefB8hUkMZ4J1UVwkzyumx2GvbVpkYUHnm",
				"tz1Qny7jVMGiwRrP9FikRK95jTNbJcffTpx1",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			key, err := NewKey(tt.input)
			testutils.CheckErr(t, tt.want.wantErr, tt.want.containsErr, err)
			assert.Equal(t, tt.want.secretKey, key.GetSecretKey())
			assert.Equal(t, tt.want.publicKey, key.PubKey.GetPublicKey())
			assert.Equal(t, tt.want.address, key.PubKey.GetPublicKeyHash())
		})
	}

}
