package main

import (
    "fmt"
    "time"

    "github.com/bwmarrin/discordgo"

    "github.com/jbillote/YAB/util"
)

/*
 * Function: RotateStatuses
 * Goroutine for rotating status messages at a regular interval
 *
 * Params:
 * session: Pointer to Discord session to update
 * statuses: Array of status messages
 * interval: Interval at which to rotate statuses, in seconds
 */
func RotateStatuses(session *discordgo.Session, statuses []string, interval int64) {
    ndx := util.RandomRange(0, len(statuses))

    err := session.UpdateGameStatus(0, statuses[ndx])
    if err != nil {
        log.Error("Unable to change status")
        log.Error(err)
    } else {
        log.Info(fmt.Sprintf("Status set to %s", statuses[ndx]))
    }

    for true {
        time.Sleep(time.Duration(interval) * time.Second)

        ndx = util.RandomRange(0, len(statuses))

        err = session.UpdateGameStatus(0, statuses[ndx])
        if err != nil {
            log.Error("Unable to change status")
            log.Error(err)
        } else {
            log.Info(fmt.Sprintf("Changed status to %s", statuses[ndx]))
        }
    }
}
