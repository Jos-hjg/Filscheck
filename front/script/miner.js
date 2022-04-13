"use strict"

$(function (){
    GetMinerInfos("f0230861")
    GetMinerInfos("f084155")
    GetMinerInfos("f061170")
    GetMinerInfos("f061174")
    GetMinerInfos("f01338337")
    GetMinerInfos("f01201012")
    GetMinerInfos("f048975")
    GetMinerInfos("f048986")
})


function GetMinerInfos(miner){
    let miner_url = "https://api2.filscout.com/api/v2/miner/" + miner
    let mining_url = "https://api2.filscout.com/api/v2/miner/" + miner + "/mining-data"
    $.ajax(miner_url,{
        type:"POST",
        success: function (res) {
            let json = {"statsType":"24h"};
            $.ajax(mining_url, {
                type: "POST",
                data:JSON.stringify(json),
                success: function (res) {
                    console.log(res)
                }
            })
            console.log(res)
        }
    })

}
