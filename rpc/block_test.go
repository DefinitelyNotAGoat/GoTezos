package rpc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Head(t *testing.T) {
	goldenBlock := getResponse(block).(*Block)
	type want struct {
		wantErr     bool
		containsErr string
		wantBlock   *Block
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(newBlockMock().handler([]byte(`not_block_data`), blankHandler)),
			want{
				true,
				"could not get head block: invalid character",
				&Block{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(newBlockMock().handler(readResponse(block), blankHandler)),
			want{
				false,
				"",
				goldenBlock,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			block, err := rpc.Head()
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.want.containsErr)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want.wantBlock, block)
		})
	}
}

func Test_Block(t *testing.T) {
	goldenBlock := getResponse(block).(*Block)
	type want struct {
		wantErr     bool
		containsErr string
		wantBlock   *Block
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(newBlockMock().handler([]byte(`not_block_data`), blankHandler)),
			want{
				true,
				"could not get block '50': invalid character",
				&Block{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(newBlockMock().handler(readResponse(block), blankHandler)),
			want{
				false,
				"",
				goldenBlock,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			block, err := rpc.Block(50)
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.wantBlock, block)
		})
	}
}

func Test_OperationHashes(t *testing.T) {
	goldenOperationHashses := getResponse(operationhashes).([][]string)

	type want struct {
		wantErr             bool
		containsErr         string
		wantOperationHashes [][]string
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(operationHashesHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"could not unmarshal operation hashes",
				[][]string{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(operationHashesHandlerMock(readResponse(operationhashes), blankHandler)),
			want{
				false,
				"",
				goldenOperationHashses,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			operationHashes, err := rpc.OperationHashes("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.wantOperationHashes, operationHashes)
		})
	}
}

func Test_BallotList(t *testing.T) {
	goldenBallotList := getResponse(ballotList).(*BallotList)

	type want struct {
		wantErr     bool
		containsErr string
		ballotList  BallotList
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(ballotListHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get ballot list",
				BallotList{},
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(ballotListHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal ballot list",
				BallotList{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(ballotListHandlerMock(readResponse(ballotList), blankHandler)),
			want{
				false,
				"",
				*goldenBallotList,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			ballotList, err := rpc.BallotList("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.ballotList, ballotList)
		})
	}
}

func Test_Ballots(t *testing.T) {
	goldenBallots := getResponse(ballots).(*Ballots)

	type want struct {
		wantErr     bool
		containsErr string
		ballots     Ballots
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(ballotsHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get ballots",
				Ballots{},
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(ballotsHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal ballots",
				Ballots{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(ballotsHandlerMock(readResponse(ballots), blankHandler)),
			want{
				false,
				"",
				*goldenBallots,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			ballots, err := rpc.Ballots("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.ballots, ballots)
		})
	}
}

func Test_CurrentPeriodKind(t *testing.T) {
	type want struct {
		wantErr           bool
		containsErr       string
		currentPeriodKind string
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(currentPeriodKindHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get current period kind",
				"",
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(currentPeriodKindHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal current period kind",
				"",
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(currentPeriodKindHandlerMock([]byte(`"promotion_vote"`), blankHandler)),
			want{
				false,
				"",
				"promotion_vote",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			currentPeriodKind, err := rpc.CurrentPeriodKind("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.currentPeriodKind, currentPeriodKind)
		})
	}
}

func Test_CurrentProposal(t *testing.T) {
	type want struct {
		wantErr         bool
		containsErr     string
		currentProposal string
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(currentProposalHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get current proposal",
				"",
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(currentProposalHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal current proposal",
				"",
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(currentProposalHandlerMock([]byte(`"promotion_vote"`), blankHandler)),
			want{
				false,
				"",
				"promotion_vote",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			currentProposal, err := rpc.CurrentProposal("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.currentProposal, currentProposal)
		})
	}
}

func Test_CurrentQuorum(t *testing.T) {
	type want struct {
		wantErr       bool
		containsErr   string
		currentQuorum int
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(currentQuorumHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get current quorum",
				0,
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(currentQuorumHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal current quorum",
				0,
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(currentQuorumHandlerMock([]byte(`7470`), blankHandler)),
			want{
				false,
				"",
				7470,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			currentQuorum, err := rpc.CurrentQuorum("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.currentQuorum, currentQuorum)
		})
	}
}

func Test_VoteListings(t *testing.T) {
	goldenVoteListings := getResponse(voteListings).(Listings)

	type want struct {
		wantErr      bool
		containsErr  string
		voteListings Listings
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(voteListingsHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get listings",
				Listings{},
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(voteListingsHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal listings",
				Listings{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(voteListingsHandlerMock(readResponse(voteListings), blankHandler)),
			want{
				false,
				"",
				goldenVoteListings,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			voteListings, err := rpc.VoteListings("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.voteListings, voteListings)
		})
	}
}

func Test_Proposals(t *testing.T) {
	goldenProposals := getResponse(proposals).(Proposals)

	type want struct {
		wantErr     bool
		containsErr string
		proposals   Proposals
	}

	cases := []struct {
		name        string
		inputHanler http.Handler
		want
	}{
		{
			"handles RPC error",
			gtGoldenHTTPMock(proposalsHandlerMock(readResponse(rpcerrors), blankHandler)),
			want{
				true,
				"failed to get proposals",
				Proposals{},
			},
		},
		{
			"failed to unmarshal",
			gtGoldenHTTPMock(proposalsHandlerMock([]byte(`junk`), blankHandler)),
			want{
				true,
				"failed to unmarshal proposals",
				Proposals{},
			},
		},
		{
			"is successful",
			gtGoldenHTTPMock(proposalsHandlerMock(readResponse(proposals), blankHandler)),
			want{
				false,
				"",
				goldenProposals,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.inputHanler)
			defer server.Close()

			rpc, err := New(server.URL)
			assert.Nil(t, err)

			proposals, err := rpc.Proposals("BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1")
			checkErr(t, tt.wantErr, tt.containsErr, err)
			assert.Equal(t, tt.want.proposals, proposals)
		})
	}
}

func Test_idToString(t *testing.T) {
	cases := []struct {
		name    string
		input   interface{}
		wantErr bool
		wantID  string
	}{
		{
			"uses integer id",
			50,
			false,
			"50",
		},
		{
			"uses string id",
			"BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1",
			false,
			"BLzGD63HA4RP8Fh5xEtvdQSMKa2WzJMZjQPNVUc4Rqy8Lh5BEY1",
		},
		{
			"uses bad id type",
			45.433,
			true,
			"",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			id, err := idToString(tt.input)
			checkErr(t, tt.wantErr, "", err)
			assert.Equal(t, tt.wantID, id)
		})
	}
}
