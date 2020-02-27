#!/bin/sh

for f in bin/jargon*; do shasum -a 256 $f > $f.sha256; done
