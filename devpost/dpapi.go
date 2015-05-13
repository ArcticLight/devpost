package main

type dpstatus struct {
    ok bool
    giterror error
    gitpath string
    
    workingdir string
    closereq chan(bool)
    firstContact bool
    usersets *usersettings
}

type usersettings struct {
    gitusercommand string
    controlprefix string
}


var status *dpstatus = &dpstatus{
    ok: false,
    giterror: nil,
    gitpath: "",
    workingdir: "",
    closereq: make(chan bool, 1),
    firstContact: true,
    usersets: &usersettings {
        gitusercommand: "",
        controlprefix: "devpost",
    },
}
