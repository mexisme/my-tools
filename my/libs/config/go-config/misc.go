package config

import (
	log "github.com/sirupsen/logrus"
	errors "golang.org/x/xerrors"

	"github.com/mexisme/go-config/settings"
)

// DryRun says whether the dry_run config has been set
func DryRun(reason string, args ...interface{}) bool {
	dryRun := settings.DryRun()
	if dryRun {
		log.Infof("DRY-RUN MODE: "+reason, args...)
	}

	return dryRun
}

// FromStringOrFunc will return a different value depending on the provided val:
// - If it's a string, provide the given val
// - If it's a func(), provide the val returned by the func
func FromStringOrFunc(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		return val.(string), nil
	case func() string:
		f := val.(func() string)
		return f(), nil
	case nil:
		return "", &MissingValue{val}
	}

	return "", &UnsupportedValue{val}
}

// MissingValueIsEmpty panics if "err" != nil or is a "MissingValue" type, otherwise it returns the "val"
func MissingValueIsEmpty(val interface{}, err error) interface{} {
	var errorAsMissing *MissingValue
	if errors.As(err, &errorAsMissing) {
		return ""
	}
	return Must(val, err)
}

// Must panics if "err" != nil, otherwise it returns the the "val"
func Must(val interface{}, err error) interface{} {
	if err != nil {
		log.Panic(err)
	}

	return val
}
