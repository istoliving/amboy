package dependency

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// AlwaysRebuildSuite tests the Always dependency implementation which
// always returns the Ready dependency state. Does contain support for
// dependency graph resolution, but the tasks will always run.
type AlwaysRebuildSuite struct {
	dep *Always
	suite.Suite
}

func TestAlwaysRebuildSuite(t *testing.T) {
	suite.Run(t, new(AlwaysRebuildSuite))
}

func (s *AlwaysRebuildSuite) SetupTest() {
	s.dep = NewAlways()
}

func (s *AlwaysRebuildSuite) TestAlwaysImplementsDependencyManagerInterface() {
	s.Implements((*Manager)(nil), s.dep)
}

func (s *AlwaysRebuildSuite) TestConstructorCreatesObjectWithExpectedValues() {
	s.True(s.dep.ShouldRebuild)
	s.Equal("always", s.dep.T.Name)
	s.Equal(0, s.dep.T.Version)
}

func (s *AlwaysRebuildSuite) TestHasComposedJobEdgesInstance() {
	s.IsType(s.dep.JobEdges, &JobEdges{})

	var ok bool
	var dep interface{} = s.dep

	_, ok = dep.(interface {
		Edges() []string
	})

	s.True(ok)

	_, ok = dep.(interface {
		AddEdge(string) error
	})

	s.True(ok)
}

func (s *AlwaysRebuildSuite) TestTypeAccessorProvidesAccessToTheCorrectTypeInfo() {
	s.Equal("always", s.dep.Type().Name)
	s.Equal(0, s.dep.Type().Version)
}

func (s *AlwaysRebuildSuite) TestAlwaysDependenciesHaveReadyState() {
	s.Exactly(Ready, s.dep.State())

	s.dep.ShouldRebuild = false
	s.Exactly(Ready, s.dep.State())

	s.dep.ShouldRebuild = true
	s.Exactly(Ready, s.dep.State())
}
