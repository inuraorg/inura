package derive

import "github.com/inuraorg/inura/inura-service/testutils"

var _ Engine = (*testutils.MockEngine)(nil)

var _ L1Fetcher = (*testutils.MockL1Source)(nil)

var _ Metrics = (*testutils.TestDerivationMetrics)(nil)
