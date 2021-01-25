target("Google-in-go")
    after_build(function (target)
        print("Add upx packet..")
        if is_plat("windows") then
            os.run("./tools/upx.exe %s", target:targetfile())
        end
        if is_plat("linux") then
            os.run("./tools/upx %s", target:targetfile())
        end
        print("Add upx end")
    end)
    set_kind("binary")
    add_ldflags("-s", "-w")
    add_files("src/*.go") 