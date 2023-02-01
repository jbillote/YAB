package graphql

import (
    "context"
    "github.com/jbillote/YAB/config"
    "github.com/jbillote/YAB/util/logger"
    "github.com/shurcooL/graphql"
)

type MBTLMove struct {
    Name      string
    Input     string
    Damage    string
    Block     string
    Cancel    string
    Property  string
    Cost      string
    Attribute string
    Startup   string
    Active    string
    Recovery  string
    Overall   string
    Advantage string
    Invuln    string
}

func QueryMBTLMove(character string, input string) (*MBTLMove, error) {
    log := logger.GetLogger("QueryMBTLMove")

    log.Info("Building GraphQL query object")

    var q struct {
        Move struct {
            Name      graphql.String
            Input     graphql.String
            Damage    graphql.String
            Block     graphql.String
            Cancel    graphql.String
            Property  graphql.String
            Cost      graphql.String
            Attribute graphql.String
            Startup   graphql.String
            Active    graphql.String
            Recovery  graphql.String
            Overall   graphql.String
            Advantage graphql.String
            Invuln    graphql.String
        } `graphql:"getMove(character: $character, input: $input)"`
    }

    queryVariables := map[string]interface{}{
        "character": graphql.String(character),
        "input":     graphql.String(input),
    }

    c := config.GetConfig()

    log.Info("Building GraphQL Query request")
    cl := graphql.NewClient(c.APIHost+"/graphql", nil)

    log.Info("Making GraphQL Query request")
    err := cl.Query(context.Background(), &q, queryVariables)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    return &MBTLMove{
        Name:      string(q.Move.Name),
        Input:     string(q.Move.Input),
        Damage:    string(q.Move.Damage),
        Block:     string(q.Move.Block),
        Cancel:    string(q.Move.Cancel),
        Property:  string(q.Move.Property),
        Cost:      string(q.Move.Cost),
        Attribute: string(q.Move.Attribute),
        Startup:   string(q.Move.Startup),
        Active:    string(q.Move.Active),
        Recovery:  string(q.Move.Recovery),
        Overall:   string(q.Move.Overall),
        Advantage: string(q.Move.Advantage),
        Invuln:    string(q.Move.Invuln),
    }, nil
}
