FROM golang:1.21 
RUN mkdir /distr
COPY ./bin/linux-amd64_deb_latest.tgz /distr
RUN cd /distr && tar -xf linux-amd64_deb_latest.tgz && cd linux-amd64_deb && ./install.sh && dpkg -i lsb-cprocsp-devel_5.0.12000-6_all.deb
RUN export GODEBUG=cgocheck=0
RUN apt update && apt install zsh -y && chsh -s $(which zsh) $(whoami) && sh -c "$(wget https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)" -y
RUN go install github.com/go-delve/delve/cmd/dlv@latest
CMD echo "This is a test." | wc -
