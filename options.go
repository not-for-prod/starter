// main details borrowed from https://github.com/twmb/franz-go/blob/master/pkg/kgo/config.go

package starter

// Opt is an option to configure a client.
type Opt interface {
	apply(*cfg)
}

type opt struct{ fn func(*cfg) }

func (opt opt) apply(cfg *cfg) { opt.fn(cfg) }

type cfg struct {
}

func (cfg *cfg) validate() error {
	return nil
}

func defaultCfg() cfg {
	return cfg{}
}

func With() Opt {
	return opt{func(_ *cfg) {

	}}
}
