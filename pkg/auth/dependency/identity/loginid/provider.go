package loginid

import (
	"sort"

	"github.com/skygeario/skygear-server/pkg/auth/dependency/identity"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/loginid"
	"github.com/skygeario/skygear-server/pkg/core/auth/metadata"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/errors"
	"github.com/skygeario/skygear-server/pkg/core/time"
	"github.com/skygeario/skygear-server/pkg/core/uuid"
)

type Provider struct {
	Store                    *Store
	Time                     time.Provider
	Config                   config.LoginIDConfiguration
	LoginIDChecker           loginid.LoginIDChecker
	LoginIDNormalizerFactory loginid.LoginIDNormalizerFactory
}

func (p *Provider) List(userID string) ([]*Identity, error) {
	is, err := p.Store.List(userID)
	if err != nil {
		return nil, err
	}

	sortIdentities(is)
	return is, nil
}

func (p *Provider) ListByClaim(name string, value string) ([]*Identity, error) {
	is, err := p.Store.ListByClaim(name, value)
	if err != nil {
		return nil, err
	}

	sortIdentities(is)
	return is, nil
}

func (p *Provider) Get(userID, id string) (*Identity, error) {
	return p.Store.Get(userID, id)
}

func (p *Provider) GetByLoginID(loginID loginid.LoginID) ([]*Identity, error) {
	im := map[string]*Identity{}
	for _, config := range p.Config.Keys {
		if !(loginID.Key == "" || config.Key == loginID.Key) {
			continue
		}

		// Normalize expects loginID is in correct type so we have to validate it first.
		invalid := p.LoginIDChecker.ValidateOne(loginid.LoginID{
			Key:   config.Key,
			Value: loginID.Value,
		})
		if invalid != nil {
			continue
		}

		normalizer := p.LoginIDNormalizerFactory.NormalizerWithLoginIDKey(config.Key)
		normalizedloginID, err := normalizer.Normalize(loginID.Value)
		if err != nil {
			return nil, err
		}

		i, err := p.Store.GetByLoginID(config.Key, normalizedloginID)
		if errors.Is(err, identity.ErrIdentityNotFound) {
			continue
		} else if err != nil {
			return nil, err
		}

		im[i.ID] = i
	}

	var is []*Identity
	for _, i := range im {
		is = append(is, i)
	}
	return is, nil
}

func (p *Provider) IsLoginIDKeyType(loginIDKey string, loginIDKeyType metadata.StandardKey) bool {
	return p.LoginIDChecker.CheckType(loginIDKey, loginIDKeyType)
}

func (p *Provider) Normalize(loginID loginid.LoginID) (normalized *loginid.LoginID, typ string, err error) {
	config := p.lookupLoginIDConfig(loginID)
	if config == nil {
		panic("loginid: unknown login ID key " + loginID.Key)
	}

	normalizer := p.LoginIDNormalizerFactory.NormalizerWithLoginIDKey(loginID.Key)
	normalizedloginID, err := normalizer.Normalize(loginID.Value)
	if err != nil {
		return
	}

	normalized = &loginid.LoginID{
		Key:   loginID.Key,
		Value: normalizedloginID,
	}
	typ = string(config.Type)
	return
}

func (p *Provider) Validate(loginIDs []loginid.LoginID) error {
	return p.LoginIDChecker.Validate(loginIDs)
}

func (p *Provider) New(userID string, loginID loginid.LoginID) *Identity {
	config := p.lookupLoginIDConfig(loginID)
	if config == nil {
		panic("loginid: unknown login ID key " + loginID.Key)
	}

	normalizer := p.LoginIDNormalizerFactory.NormalizerWithLoginIDKey(loginID.Key)
	normalizedloginIDValue, err := normalizer.Normalize(loginID.Value)
	if err != nil {
		panic(errors.Newf("loginid: failed to normalize login ID: %w", err))
	}
	uniqueKey, err := normalizer.ComputeUniqueKey(normalizedloginIDValue)
	if err != nil {
		panic(errors.Newf("loginid: failed to compute login ID unique key: %w", err))
	}

	claims := map[string]string{}
	if standardKey, ok := p.LoginIDChecker.StandardKey(loginID.Key); ok {
		claims[string(standardKey)] = normalizedloginIDValue
	}

	return &Identity{
		ID:              uuid.New(),
		UserID:          userID,
		LoginIDKey:      loginID.Key,
		LoginID:         normalizedloginIDValue,
		OriginalLoginID: loginID.Value,
		UniqueKey:       uniqueKey,
		Claims:          claims,
	}
}

func (p *Provider) Create(i *Identity) error {
	return p.Store.Create(i)
}

func (p *Provider) Delete(i *Identity) error {
	return p.Store.Delete(i)
}

func (p *Provider) lookupLoginIDConfig(loginID loginid.LoginID) *config.LoginIDKeyConfiguration {
	for _, c := range p.Config.Keys {
		if c.Key == loginID.Key {
			return &c
		}
	}
	return nil
}

func sortIdentities(is []*Identity) {
	sort.Slice(is, func(i, j int) bool {
		return is[i].UniqueKey < is[j].UniqueKey
	})
}