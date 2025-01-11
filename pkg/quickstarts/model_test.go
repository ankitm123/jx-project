//go:build unit
// +build unit

package quickstarts_test

import (
	"testing"

	"github.com/jenkins-x-plugins/jx-project/pkg/quickstarts"
	"github.com/stretchr/testify/assert"
)

func TestquickstartsQuickstartModelFilterText(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/ruby",
		Name: "ruby",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ruby"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		Text: "ruby",
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 1, len(results))
	assert.Contains(t, results, quickstart3)
}

func TestquickstartsQuickstartModelFilterTextMatchesMoreThanOne(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/ruby",
		Name: "ruby",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ruby"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		Text: "node-htt",
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, quickstart1)
	assert.Contains(t, results, quickstart2)
}

func TestquickstartsQuickstartModelFilterTextMatchesOneExactly(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/ruby",
		Name: "ruby",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ruby"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		Text: "node-http",
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 1, len(results))
	assert.Contains(t, results, quickstart1)
}

func TestquickstartsQuickstartModelFilterExcludesMachineLearning(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "machine-learning-quickstarts/ML-is-a-machine-learning-quickstart",
		Name: "ML-is-a-machine-learning-quickstart",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ML-is-a-machine-learning-quickstart"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		AllowML: false,
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, quickstart1)
	assert.Contains(t, results, quickstart2)
	assert.NotContains(t, results, quickstart3)
}

func TestquickstartsQuickstartModelFilterIncludesMachineLearning(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "machine-learning-quickstarts/ML-is-a-machine-learning-quickstart",
		Name: "ML-is-a-machine-learning-quickstart",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ML-is-a-machine-learning-quickstart"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		AllowML: true,
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 3, len(results))
	assert.Contains(t, results, quickstart1)
	assert.Contains(t, results, quickstart2)
	assert.Contains(t, results, quickstart3)
}

func TestquickstartsQuickstartModelFilterDefaultsToNoMachineLearning(t *testing.T) {
	t.Parallel()

	quickstart1 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http",
		Name: "node-http",
	}
	quickstart2 := &quickstarts.Quickstart{
		ID:   "jenkins-x-quickstarts/node-http-watch-pipeline-activity",
		Name: "node-http-watch-pipeline-activity",
	}
	quickstart3 := &quickstarts.Quickstart{
		ID:   "machine-learning-quickstarts/ML-is-a-machine-learning-quickstart",
		Name: "ML-is-a-machine-learning-quickstart",
	}

	qstarts := make(map[string]*quickstarts.Quickstart)
	qstarts["node-http"] = quickstart1
	qstarts["node-http-watch-pipeline-activity"] = quickstart2
	qstarts["ML-is-a-machine-learning-quickstart"] = quickstart3

	quickstartModel := &quickstarts.QuickstartModel{
		Quickstarts: qstarts,
	}

	quickstartFilter := &quickstarts.QuickstartFilter{
		Text: "",
	}

	results := quickstartModel.Filter(quickstartFilter)

	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, quickstart1)
	assert.Contains(t, results, quickstart2)
	assert.NotContains(t, results, quickstart3)
}
