"use strict"

$(function (){
    // $('#miner-list div').empty()
    GetMinerInfos()
    // GetMinerInfos("f084155")
    // GetMinerInfos("f061170")
    // GetMinerInfos("f061174")
    // GetMinerInfos("f01338337")
    // GetMinerInfos("f01201012")
    // GetMinerInfos("f048975")
    // GetMinerInfos("f048986")
})




function GetMinerInfos(){
    // let miner_url = "https://api2.filscout.com/api/v2/miner/" + miner
    // let mining_url = "https://api2.filscout.com/api/v2/miner/" + miner + "/mining-data"

    let miner_url = "http://localhost:8082/miner"
    let mining_url = "http://localhost:8082/mining-data"
    $.ajax(miner_url,{
        type:"GET",
        success: function (miner_res) {
            let json = {"statsType":"24h"};
            $.ajax(mining_url, {
                type: "GET",
                data:JSON.stringify(json),
                success: function (mining_res) {

                    let str = ""
                    for (let miner of miner_res.data){
                        for (let mining of mining_res.data){
                            if (miner["miner"] === mining["miner"]){
                                mining['luckyValue'] = (mining['luckyValue'] * 100).toFixed(2) + "%"
                                let container = `
                    <div class="contain" data-v-5d5bb70e>
                        <div class="account-tip" data-v-5d5bb70e data-v-574bbb96>
                            <div class="account-title" data-v-5d5bb70e data-v-574bbb96>
                                <span class="account-name" data-v-5d5bb70e data-v-574bbb96>账户 ${miner.miner}</span>
                            </div>
                        </div>
                        <div class="item" data-v-5d5bb70e data-v-574bbb96>
                            <div data-v-b4d9e120 data-v-574bbb96>
                                <div class="card-main card-radius" data-v-d66790d0 data-v-b4d9e120>
                                    <div class="card-top card-title-radius" data-v-d66790d0>
                                        <span data-v-d66790d0 data-v-b4d9e120>算力概览</span></div>
                                    <div class="card-contain" data-v-d66790d0>
                                        <div class="power-contain" data-v-d66790d0 data-v-b4d9e120>
                                            <div data-v-d66790d0 data-v-b4d9e120>
                                                <div data-v-d66790d0 data-v-b4d9e120>
                                                    <div class="power_sl" data-v-d66790d0 data-v-b4d9e120>有效算力</div>
                                                    <div class="total" data-v-d66790d0 data-v-b4d9e120>${miner.qualityPowerStr}</div>
                                                </div>
                                                <div class="total-persent" data-v-d66790d0 data-v-b4d9e120>
                                                    <div data-v-d66790d0 data-v-b4d9e120>
                                                        <span data-v-d66790d0 data-v-b4d9e120>算力占比:</span>
                                                        <span data-v-d66790d0 data-v-b4d9e120>${miner.qualityPowerPercentStr}</span>
                                                    </div>
                                                </div>
                                                <div class="total-persent" data-v-d66790d0 data-v-b4d9e120>
                                                    <div data-v-d66790d0 data-v-b4d9e120>
                                                        <span data-v-d66790d0 data-v-b4d9e120>排名:</span>
                                                        <span data-v-d66790d0 data-v-b4d9e120>${miner.powerRank}</span></div>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-b4d9e120>
                                                <div data-v-d66790d0 data-v-b4d9e120>
                                                    <span data-v-d66790d0 data-v-b4d9e120>原值算力:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-b4d9e120>${miner.qualityPowerStr}</span>
                                                </div>
                                                <div class="miner_desc" data-v-d66790d0 data-v-b4d9e120>
                                                    <span data-v-d66790d0 data-v-b4d9e120>累计出块份数:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-b4d9e120>${miner.winCount}</span>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-b4d9e120>
                                                <div data-v-d66790d0 data-v-b4d9e120>
                                                    <span data-v-d66790d0 data-v-b4d9e120>累计出块奖励:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-b4d9e120>${miner.blockRewardStr}</span>
                                                </div>
                                                <div data-v-d66790d0 data-v-b4d9e120>
                                                    <span data-v-d66790d0 data-v-b4d9e120>扇区大小:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-b4d9e120>${miner.sectorSizeStr}</span>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-b4d9e120>
                                                <div data-v-d66790d0 data-v-b4d9e120>
                                                    <span data-v-d66790d0 data-v-b4d9e120>扇区状态:</span>
                                                    <span class="status-all" data-v-d66790d0 data-v-b4d9e120> ${miner.sectorCount} 全部,</span>
                                                    <span class="status-success" data-v-d66790d0 data-v-b4d9e120> ${miner.activeCount} 有效,</span>
                                                    <span class="status-fail" data-v-d66790d0 data-v-b4d9e120> ${miner.faultCount} 错误,</span>
                                                    <span class="status-warn" data-v-d66790d0 data-v-b4d9e120> ${miner.recoveryCount} 恢复中</span>
                                                </div>
                                            </div>

                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div data-v-9eb4ae2e data-v-574bbb96>
                                <div class="card-main card-radius" data-v-d66790d0 data-v-9eb4ae2e>
                                    <div class="card-top card-title-radius" data-v-d66790d0>
                                        <div class="mining-title" data-v-d66790d0 data-v-9eb4ae2e>
                                            <div class="more_click" data-v-d66790d0 data-v-9eb4ae2e>
                                                <span data-v-d66790d0 data-v-9eb4ae2e>产出统计</span></div>
                                            <div class="time-main" data-v-36ae89d1 data-v-9eb4ae2e>
                                                <span class="item selected" data-v-36ae89d1>24时</span><span class="item" data-v-36ae89d1>7天</span><span class="item" data-v-36ae89d1>30天</span>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="card-contain" data-v-d66790d0>
                                        <div class="mining-contain" data-v-d66790d0 data-v-9eb4ae2e>
                                            <div data-v-d66790d0 data-v-9eb4ae2e>
                                                <div data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>算力增量:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.qualityPowerGrowthStr}</span>
                                                </div>
                                                <div data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>算力增速:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.provingPowerStr}</span>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-9eb4ae2e>
                                                <div class="miner_desc" data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>节点当量:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.machinesNum}</span>
                                                </div>
                                                <div class="miner_desc" data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>出块数量:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.blocks}</span>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-9eb4ae2e>
                                                <div data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>出块奖励:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.blockRewardStr}</span>
                                                </div>
                                                <div data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>出块奖励占比:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.blockRewardPercent}</span>
                                                </div>
                                            </div>
                                            <div data-v-d66790d0 data-v-9eb4ae2e>
                                                <div class="miner_desc" data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>产出效率:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.miningEfficiency}</span>
                                                </div>
                                                <div class="miner_desc" data-v-d66790d0 data-v-9eb4ae2e>
                                                    <span data-v-d66790d0 data-v-9eb4ae2e>幸运值:</span>
                                                    <span class="amount" data-v-d66790d0 data-v-9eb4ae2e>${mining.luckyValue}</span>
                                                </div>
                                            </div>
                                            <div style="height:48px;" data-v-d66790d0 data-v-b4d9e120></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    
`
                                str += container
                            }
                        }

                    }
                    $('#miner-list div').empty().append(str)
                    // console.log(str)

                }
            })

        }
    })

}
