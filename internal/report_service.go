package internal

import "io"

type Service struct {
	dependenciesAdapter DependenciesAdapter
	reporterAdapter     ReporterAdapter
}

func NewService(depenciesAdapter DependenciesAdapter, reporterAdapter ReporterAdapter) *Service {
	return &Service{
		dependenciesAdapter: depenciesAdapter,
		reporterAdapter:     reporterAdapter,
	}
}

func (s *Service) GetReport(r1 io.Reader, r2 io.Reader) error {
	diff, err := s.dependenciesAdapter.GetDiff(r1, r2)
	if err != nil {
		return err
	}

	err = s.reporterAdapter.Print(diff)
	if err != nil {
		return err
	}

	return nil
}
