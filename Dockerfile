FROM golang:1.7.1-alpine
MAINTAINER "Team Slizco" <teamslizco@gmail.com>

ARG git_url
ARG git_branch_name
ARG git_sha1
ARG project_name
ARG project_go_version

EXPOSE 8080

COPY pkg/linux_amd64/recorder /usr/bin/recorder

LABEL slizco.git.url=${git_url} \
      slizco.git.branch-name=${git_branch_name} \
      slizco.git.sha1=${git_sha1} \
      slizco.project.name=${project_name} \
      slizco.project.version=${project_go_version}

CMD ["/usr/bin/recorder"]
