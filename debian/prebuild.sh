#!/bin/sh
GIT_VERSION=$(git -C $GBP_GIT_DIR rev-parse HEAD)
TAG_VERSION=$(git -C $GBP_GIT_DIR tag -l --contains $GIT_VERSION)

$GBP_BUILD_DIR/gen_version_file.sh $GIT_VERSION $TAG_VERSION

cp $GBP_BUILD_DIR/data/coyim.gschema.override $GBP_BUILD_DIR/gui/settings/definitions/
make -C $GBP_BUILD_DIR/gui/settings/definitions generate
