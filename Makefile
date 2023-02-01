ROOT = main.go
PARSE_MESSAGE = parse_message.go
ROTATE_STATUS = rotate_status.go
COMMANDS := $(wildcard commands/*)
CONFIG := $(wildcard config/*)
FRAMEDATA := $(wildcard framedata/*)
GRAPHQL := $(wildcard graphql/*)
TWITTER := $(wildcard twitter/*)
UTIL := $(wildcard util/*)

toukabot: $(ROOT) $(PARSE_MESSAGE) $(ROTATE_STATUS) $(COMMANDS) $(CONFIG) $(FRAMEDATA) $(GRAPHQL) $(TWITTER) $(UTIL)
	go build

clean:
	rm toukabot-v2