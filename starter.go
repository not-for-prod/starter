package starter

type Starter struct {
	cfg cfg
}

func New(opts ...Opt) (*Starter, error) {
	config := defaultCfg()

	for _, opt := range opts {
		opt.apply(&config)
	}

	err := config.validate()
	if err != nil {
		return nil, err
	}

	return &Starter{cfg: config}, nil
}

func (a *Starter) Run(_ ...Service) {

}
