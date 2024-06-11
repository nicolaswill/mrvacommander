package server

import (
	"mrvacommander/pkg/common"
	"mrvacommander/pkg/logger"
	"mrvacommander/pkg/queue"
	"mrvacommander/pkg/storage"
)

type SessionInfo struct {
	ID             int
	Owner          string
	ControllerRepo string

	QueryPack    string
	Language     string
	Repositories []common.OwnerRepo

	AccessMismatchRepos []common.OwnerRepo
	NotFoundRepos       []common.OwnerRepo
	NoCodeqlDBRepos     []common.OwnerRepo
	OverLimitRepos      []common.OwnerRepo

	AnalysisRepos *map[common.OwnerRepo]storage.DBLocation
}

type CommanderSingle struct {
	st *CommanderVisibles
}

func NewCommanderSingle() *CommanderSingle {
	c := CommanderSingle{}
	return &c
}

// type State struct {
// 	Commander Commander
// 	Logger    logger.Logger
// 	Queue     queue.Queue
// 	Storage   storage.Storage
// 	Runner    agent.Runner
// }

type CommanderVisibles struct {
	Logger      logger.Logger
	Queue       queue.Queue
	ServerStore storage.Storage
	// TODO extra package for query pack storage
	QueryPackStore storage.Storage
	// TODO extra package for ql db storage
	QLDBStore storage.Storage
}
