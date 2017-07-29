// This Jenkinsfile is a groovy script that is executed top to bottom during CI.
//
// The Jenkinsfile build flow allows using the programming language to configure
// a build with custom workflows. In the case of dataseries, the primary use case
// is to build bare-bones Go binary images instead of using the 900+MB go builder
// image on deploy.
//
// See https://underarmour.atlassian.net/wiki/display/INFRA/%5BBuild+Flow%5D+MultiBranch+Workflow
// for more info.


/////////// DEFINITIONS
// Here we define arrays describing all our docker builds and tests.

// The build image is where we build all our go binaries.
// In order to allow the release to work appropriately, we build this before all the others.
def precursor_images = [
        build: [
                DOCKER_IMAGE_NAME: "svc_friends/build",
                DOCKER_CONTEXT_PATH: ".",
                DOCKER_FILE: "docker/build/Dockerfile",
        ],
]

// Images used for support of tests and localdev.
def support_images = [
/*    goldengate: [
        DOCKER_IMAGE_NAME: "svc_friends/goldengate",
        DOCKER_CONTEXT_PATH: "docker/goldengate",
        DOCKER_FILE: "docker/goldengate/Dockerfile",
    ],*/
    sentry: [
            DOCKER_IMAGE_NAME: "svc_friends/sentry",
            DOCKER_CONTEXT_PATH: "docker/sentry",
            DOCKER_FILE: "docker/sentry/Dockerfile",
    ],
]

// Images used in production. Will be released on tags.
def prod_images = [
        // friendsapi uses the PREBUILD_SCRIPT function
        // to run a script to prepare the workspace -before- running a docker build.
        friendsapi: [
                DOCKER_IMAGE_NAME: "svc_friends/friendsapi",
                DOCKER_CONTEXT_PATH: ".",
                DOCKER_FILE: "docker/friendsapi/Dockerfile",
                PREBUILD_SCRIPT: "docker/build/extract-binaries.sh"
        ],
]

// Docker-based test suites are defined here.
def tests = [
        functional: [
                TEST_YAML: "./test/local_functional.yml",
                TEST_TARGET: "test",
                HAS_JUNIT_REPORT: true,
        ],
]

/////////// LIBRARY INITIALIZATION
// Here we download and initialize the infra/CI library which we will use for workflow.

env.GIT_PROJECT = "uarecord/svc_friends"

// EDIT the following to alert a slack room when builds complete
// env.SLACK_ROOM_FINISH = "<MY_TEAM>-builds"

stage 'Init'
gitlab_v1 = node('swarm') {
    git url: 'git@scm-main-01.dc.myfitnesspal.com:infra/CI.git', branch: 'master'
    load 'pipelines/gitlab_v1.groovy'
}

gitlab_v1.pipeline_map([
        precursor_images: precursor_images + support_images,
        images: prod_images,
        tests: tests,
])
