#!/bin/bash

RESTIC_PASSWORD="<my password>" /Users/armhold/work/go/bin/restic -r /Volumes/backup1/x/restic-repo-testing backup ~ --host george-iMac --exclude-file ~/restic-excludes
