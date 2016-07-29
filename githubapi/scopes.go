package githubapi

import (
	"bytes"
	"strings"
)

// Scopes ... type used to represent github scopes used with oauth.
//(no scope)  Grants read-only access to public information (includes public user profile info, public repository info, and gists)
//user    Grants read/write access to profile info only. Note that this scope includes user:email and user:follow.
//user:email  Grants read access to a user's email addresses.
//user:follow Grants access to follow or unfollow other users.
//public_repo Grants read/write access to code, commit statuses, collaborators, and deployment statuses for public repositories and organizations. Also required for starring public repositories.
//repo    Grants read/write access to code, commit statuses, repository invitations, collaborators, and deployment statuses for public and private repositories and organizations.
//repo_deployment Grants access to deployment statuses for public and private repositories. This scope is only necessary to grant other users or services access to deployment statuses, without granting access to the code.
//repo:status Grants read/write access to public and private repository commit statuses. This scope is only necessary to grant other users or services access to private repository commit statuses without granting access to the code.
//delete_repo Grants access to delete adminable repositories.
//notifications   Grants read access to a user's notifications. repo also provides this access.
//gist    Grants write access to gists.
//read:repo_hook  Grants read and ping access to hooks in public or private repositories.
//write:repo_hook Grants read, write, and ping access to hooks in public or private repositories.
//admin:repo_hook Grants read, write, ping, and delete access to hooks in public or private repositories.
//admin:org_hook  Grants read, write, ping, and delete access to organization hooks. Note: OAuth tokens will only be able to perform these actions on organization hooks which were created by the OAuth application. Personal access tokens will only be able to perform these actions on organization hooks created by a user.
//read:org    Read-only access to organization, teams, and membership.
//write:org   Publicize and unpublicize organization membership.
//admin:org   Fully manage organization, teams, and memberships.
//read:public_key List and view details for public keys.
//write:public_key    Create, list, and view details for public keys.
//admin:public_key    Fully manage public keys.
//read:gpg_key    List and view details for GPG keys.
//write:gpg_key   Create, list, and view details for GPG keys.
//admin:gpg_key Fully manage GPG keys.
type Scopes struct {
	storedScopes map[string]bool
}

func (s *Scopes) fromString(scopes string) {
	s.storedScopes = make(map[string]bool)
	ss := strings.Split(scopes, ",")
	for _, scope := range ss {
		s.storedScopes[scope] = true
	}
}

func (s *Scopes) String() string {
	var buffer bytes.Buffer
	s.RepoStatus().AdminRepoHook().UserEmail().Repo().AdminOrgHook().AdminRepoHook()
	for k, allowed := range s.storedScopes {
		if allowed {
			buffer.WriteString(k)
			buffer.WriteString(" ")
		}
	}
	ss := buffer.String()
	return ss
}

// User See Comment for Scopes Type
func (s *Scopes) User() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["user"] = true
	return s
}

// UserEmail See Comment for Scopes Type
func (s *Scopes) UserEmail() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["user:email"] = true
	return s
}

// UserFollow See Comment for Scopes Type
func (s *Scopes) UserFollow() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["user:follow"] = true
	return s
}

// PublicRepo See Comment for Scopes Type
func (s *Scopes) PublicRepo() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["public_repo"] = true
	return s
}

// Repo See Comment for Scopes Type
func (s *Scopes) Repo() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["repo"] = true
	return s
}

// RepoDeployment See Comment for Scopes Type
func (s *Scopes) RepoDeployment() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["repo_deployment"] = true
	return s
}

// RepoStatus See Comment for Scopes Type
func (s *Scopes) RepoStatus() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["repo:status"] = true
	return s
}

// DeleteRepo See Comment for Scopes Type
func (s *Scopes) DeleteRepo() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["delete_repo"] = true
	return s
}

// Notifications See Comment for Scopes Type
func (s *Scopes) Notifications() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["notifications"] = true
	return s
}

// Gist See Comment for Scopes Type
func (s *Scopes) Gist() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["gist"] = true
	return s
}

// ReadRepoHook See Comment for Scopes Type
func (s *Scopes) ReadRepoHook() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["read:repo_hook"] = true
	return s
}

// WriteRepoHook See Comment for Scopes Type
func (s *Scopes) WriteRepoHook() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["write:repo_hook"] = true
	return s
}

// AdminRepoHook See Comment for Scopes Type
func (s *Scopes) AdminRepoHook() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["admin:repo_hook"] = true
	return s
}

// AdminOrgHook See Comment for Scopes Type
func (s *Scopes) AdminOrgHook() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["admin:org_hook"] = true
	return s
}

// ReadOrg See Comment for Scopes Type
func (s *Scopes) ReadOrg() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["read:org"] = true
	return s
}

// WriteOrg See Comment for Scopes Type
func (s *Scopes) WriteOrg() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["write:org"] = true
	return s
}

// AdminOrg See Comment for Scopes Type
func (s *Scopes) AdminOrg() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["admin:org"] = true
	return s
}

// ReadPublicKey See Comment for Scopes Type
func (s *Scopes) ReadPublicKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["read:public_key"] = true
	return s
}

// WritePublicKey See Comment for Scopes Type
func (s *Scopes) WritePublicKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["write:public_key"] = true
	return s
}

// AdminPublicKey See Comment for Scopes Type
func (s *Scopes) AdminPublicKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["admin:public_key"] = true
	return s
}

// ReadGpgKey See Comment for Scopes Type
func (s *Scopes) ReadGpgKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["read:gpg_key"] = true
	return s
}

// WriteGpgKey See Comment for Scopes Type
func (s *Scopes) WriteGpgKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["write:gpg_key"] = true
	return s
}

// AdminGpgKey See Comment for Scopes Type
func (s *Scopes) AdminGpgKey() *Scopes {
	if s.storedScopes == nil {
		s.storedScopes = make(map[string]bool)
	}
	s.storedScopes["admin:gpg_key"] = true
	return s
}
