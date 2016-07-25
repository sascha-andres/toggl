FROM scratch 
MAINTAINER Sascha Andres <sascha.andres@outlook.com> 
 
ADD toggl toggl
ENTRYPOINT [ "/toggl" ]