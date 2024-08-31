package consensus

// PBFT pbft 的实现接口
type PBFT interface {
	StartConsensus(request *RequestMsg) (*PrePrepareMsg, error) // 开始一个共识
	PrePrepare(prePrepareMsg *PrePrepareMsg) (*VoteMsg, error)  // 预准备阶段
	Prepare(prepareMsg *VoteMsg) (*VoteMsg, error)              // 准备阶段
	Commit(commitMsg *VoteMsg) (*ReplyMsg, *RequestMsg, error)  // 提交阶段
}
