package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	host2 "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"net/http"
)

type memory struct {
	Total       uint64  `json:"total,omitempty"`
	Free        uint64  `json:"free,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty"`
}

func HomeData(ctx *gin.Context) {
	mp := make(map[string]interface{})

	v, _ := mem.VirtualMemory()

	mes := make([]*memory, 1)

	me := &memory{
		Total:       v.Total,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
	}
	mes[0] = me
	mp["memory"] = mes

	cpu, _ := cpu.Info()
	mp["cpu"] = cpu

	hosts := make([]*host2.InfoStat, 1)
	host, _ := host2.Info()
	hosts[0] = host
	mp["host"] = hosts

	ctx.JSON(http.StatusOK, mp)
}
