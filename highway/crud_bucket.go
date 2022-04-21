package highway

import (
	context "context"
	"log"

	"github.com/gin-gonic/gin"
	btt "github.com/sonr-io/blockchain/x/bucket/types"
	bt "go.buf.build/sonr-io/grpc-gateway/sonr-io/blockchain/bucket"
)

// CreateBucket creates a new bucket.
func (s *HighwayServer) CreateBucket(ctx context.Context, req *bt.MsgCreateBucket) (*bt.MsgCreateBucketResponse, error) {
	resp, err := s.cosmos.BroadcastCreateBucket(btt.NewMsgCreateBucketFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &bt.MsgCreateBucketResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhichIs: btt.NewWhichIsToBuf(resp.WhichIs),
	}, nil
}

// CreateBucketHTTP creates a new bucket via HTTP.
func (s *HighwayServer) CreateBucketHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req bt.MsgCreateBucket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create the bucket
	resp, err := s.grpcClient.CreateBucket(s.ctx, &req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(200, gin.H{
		"code":     resp.Code,
		"message":  resp.Message,
		"which_is": btt.NewWhichIsFromBuf(resp.WhichIs),
	})
}

// UpdateBucket updates a bucket.
func (s *HighwayServer) UpdateBucket(ctx context.Context, req *bt.MsgUpdateBucket) (*bt.MsgUpdateBucketResponse, error) {
	resp, err := s.cosmos.BroadcastUpdateBucket(btt.NewMsgUpdateBucketFromBuf(req))
	if err != nil {
		return nil, err
	}
	log.Println(resp.String())
	return &bt.MsgUpdateBucketResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhichIs: btt.NewWhichIsToBuf(resp.WhichIs),
	}, nil
}

// UpdateBucketHTTP updates a bucket via HTTP.
func (s *HighwayServer) UpdateBucketHTTP(c *gin.Context) {
	// Unmarshal the request body
	var req bt.MsgUpdateBucket
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update the bucket
	resp, err := s.grpcClient.UpdateBucket(s.ctx, &req)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	// Return the response
	c.JSON(200, gin.H{
		"code":     resp.Code,
		"message":  resp.Message,
		"which_is": btt.NewWhichIsFromBuf(resp.WhichIs),
	})
}
