package metrics

import (
	"github.com/inuraorg/inura/inura-service/eth"
	opmetrics "github.com/inuraorg/inura/inura-service/metrics"
	txmetrics "github.com/inuraorg/inura/inura-service/txmgr/metrics"
)

type noopMetrics struct {
	opmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordL2BlocksProposed(l2ref eth.L2BlockRef) {}
