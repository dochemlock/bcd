// THIS FILE IS AUTO-GENERATED DO NOT EDIT

package rtorrent

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bytesizedhosting/bcd/jobs"
	"github.com/bytesizedhosting/bcd/plugins"
)

type RtorrentRPC struct {
	base *Rtorrent
	plugins.BaseRPC
}
func (self *RtorrentRPC) Reinstall(opts *RtorrentOpts, job *jobs.Job) error {
	err := self.base.Uninstall(&plugins.AppConfig{ContainerId: opts.ContainerId})
	if err != nil {
		log.Infoln("Could not remove Docker container but since this is a reinstall we don't care.")
	}
	self.Install(opts, job)
	return nil
}

func (self *RtorrentRPC) Install(opts *RtorrentOpts, job *jobs.Job) error {
	*job = *jobs.New(opts)
	log.Debugln("Rtorrent options:", opts)
	go func() {
		err := self.base.Install(opts)
		job.Options = *opts

		if err != nil {
			log.Debugln("Rtorrent installation received an error:", err)
			job.ErrorString = err.Error()
			job.Error = err
			job.Status = jobs.FAILED
		} else {
			log.Infoln("Rtorrent installation completed")
			job.Status = jobs.FINISHED
		}

		jobs.Storage.Set(job.Id, job)
	}()

	return nil
}
