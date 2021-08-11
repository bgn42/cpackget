/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright Contributors to the cpackget project. */

package installer

import (
	log "github.com/sirupsen/logrus"

	"github.com/open-cmsis-pack/cpackget/cmd/xml"
)

// PdscType is the struct that represents the installation of a
// pack via PDSC file
type PdscType struct {
	xml.PdscTag

	// file points to the actual PDSC file
	file *xml.PdscXML

	// isInstalled tells whether the pack is already installed
	isInstalled bool

	// path points to a file in the local system, whether or not it's local
	path string
}

// toPdscTag returns a <pdsc> tag representation of this PDSC file
func (p *PdscType) toPdscTag() (xml.PdscTag, error) {
	tag := p.PdscTag

	if p.file == nil {
		p.file = xml.NewPdsc(p.path)
		if err := p.file.Read(); err != nil {
			return tag, err
		}
	}

	// uses the Version from the actual file
	tag.Version = p.file.Tag().Version

	return tag, nil
}

// install installs a pack via PDSC file.
// It:
//   - Adds it to the "CMSIS_PACK_ROOT/.Local/local_repository.pidx"
//     using version from the PDSC file
func (p *PdscType) install(installation *PacksInstallationType) error {
	log.Debugf("Installing \"%s\"", p.path)
	tag, err := p.toPdscTag()
	if err != nil {
		return err
	}

	if err := installation.localPidx.AddPdsc(tag); err != nil {
		return err
	}

	return nil
}