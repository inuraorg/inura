package op_e2e

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// BuildOpProgramClient builds the `inura-program` client executable and returns the path to the resulting executable
func BuildOpProgramClient(t *testing.T) string {
	t.Log("Building inura-program-client")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "make", "inura-program-client")
	cmd.Dir = "../inura-program"
	cmd.Stdout = os.Stdout // for debugging
	cmd.Stderr = os.Stderr // for debugging
	require.NoError(t, cmd.Run(), "Failed to build inura-program-client")
	t.Log("Built inura-program-client successfully")
	return "../inura-program/bin/inura-program-client"
}
