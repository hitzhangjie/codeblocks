// +build !windows

package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsWritable(t *testing.T) {
	// root用户不需要测试
	if os.Geteuid() == 0 {
		t.SkipNow()
	}

	wd := filepath.Join(os.TempDir(), "fs_test")
	os.MkdirAll(wd, os.ModePerm)
	defer os.RemoveAll(wd)

	os.MkdirAll(filepath.Join(wd, "perm"), 0777)
	os.OpenFile(filepath.Join(wd, "perm/a"), os.O_CREATE, 0777)
	os.MkdirAll(filepath.Join(wd, "perm/b"), 0444)
	os.MkdirAll(filepath.Join(wd, "perm/c"), 0777)

	tests := []struct {
		name           string
		path           string
		wantIsWritable bool
		wantErr        bool
	}{
		{
			name:           "1-not existed dir",
			path:           "/xxxxxx/yyyyyy/zzzzz",
			wantIsWritable: false,
			wantErr:        true,
		},
		{
			name:           "2-not directory",
			path:           filepath.Join(wd, "perm/a"),
			wantIsWritable: false,
			wantErr:        true,
		},
		{
			name:           "3-not writable",
			path:           filepath.Join(wd, "perm/b"),
			wantIsWritable: false,
			wantErr:        false,
		},
		{
			name:           "4-is writable",
			path:           filepath.Join(wd, "perm/c"),
			wantIsWritable: true,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsWritable, err := IsWritable(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsWritable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsWritable != tt.wantIsWritable {
				t.Errorf("IsWritable() gotIsWritable = %v, want %v", gotIsWritable, tt.wantIsWritable)
			}
		})
	}
}
