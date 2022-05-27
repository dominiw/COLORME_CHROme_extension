
all: build

########################################################################
##                             GOLANG                                 ##
########################################################################

# If GOPATH isn't defined then set its default location.
ifeq (,$(strip $(GOPATH)))
GOPATH := $(HOME)/go
else
# If GOPATH is already set then update GOPATH to be its own
# first element.
GOPATH := $(word 1,$(subst :, ,$(GOPATH)))