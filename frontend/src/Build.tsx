import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getBuild, type buildData } from "./api/getBuild";
import { timeAgo } from "./helpers/timeAgo";

async function handleBuild(
    id: string,
    set: React.Dispatch<React.SetStateAction<buildData | undefined>>,
) {
    const build = await getBuild(id);
    if (!build) {
        return;
    }
    set(build);
}

function Build() {
    const navigate = useNavigate();
    const { id } = useParams();

    const [build, setBuild] = useState<buildData | undefined>();

    const buildClass =
        build?.Data.Build.AscendClassName == "None"
            ? build.Data.Build.ClassName
            : build?.Data.Build.AscendClassName;
    const title = `Level ${build?.Data.Build.Level} ${buildClass}`;
    document.title = title;

    console.log(build);

    useEffect(() => {
        if (!id) {
            navigate("/");
            return;
        }

        const init = async () => {
            await handleBuild(id, setBuild);
        };
        init();
    }, []);

    return (
        <div className="min-h-screen text-zinc-100">
            <h1
                className="text-2xl font-semibold mb-4 mt-4 text-slate-100 cursor-pointer"
                onClick={() => navigate("/")}
            >
                poeb.in
            </h1>
            <div className="max-w-7xl mx-auto grid grid-cols-12 gap-6 px-6 py-6">
                {/* Sidebar */}
                <aside className="col-span-3 space-y-6">
                    <div className="rounded-2xl bg-slate-900 border border-slate-400 p-6 shadow-lg">
                        <h2 className="text-lg font-semibold mb-4">
                            Build Info
                        </h2>
                        <div className="space-y-2 text-sm text-zinc-400">
                            <div>
                                <span className="text-zinc-200">Class:</span>{" "}
                                {build?.Data.Build.ClassName}
                            </div>
                            <div>
                                <span className="text-zinc-200">
                                    Ascendancy:
                                </span>{" "}
                                {build?.Data.Build.AscendClassName}
                            </div>
                            <div>
                                <span className="text-zinc-200">Level: </span>
                                {build?.Data.Build.Level}
                            </div>
                            <div>
                                {build ? timeAgo(build?.LastModified) : null}
                            </div>
                        </div>
                    </div>

                    <div className="rounded-2xl bg-slate-900 border border-slate-400 p-6 shadow-lg">
                        <h2 className="text-lg font-semibold mb-4">Defences</h2>
                        <div className="grid grid-cols-2 gap-3 text-sm">
                            {[
                                [
                                    "Life",
                                    build?.Data.Build.Stats.find(
                                        (e: { Name: string; value: string }) =>
                                            e.Name === "Life",
                                    ).Value,
                                ],
                                [
                                    "ES",
                                    build?.Data.Build.Stats.find(
                                        (e: { Name: string; value: string }) =>
                                            e.Name === "EnergyShield",
                                    ).Value,
                                ],
                                [
                                    "Armour",
                                    build?.Data.Build.Stats.find(
                                        (e: { Name: string; value: string }) =>
                                            e.Name === "Armour",
                                    ).Value,
                                ],
                                [
                                    "Evasion",
                                    build?.Data.Build.Stats.find(
                                        (e: { Name: string; value: string }) =>
                                            e.Name === "Evasion",
                                    ).Value,
                                ],
                            ].map(([label, value]) => (
                                <div
                                    key={label}
                                    className="bg-slate-800 p-3 rounded-xl"
                                >
                                    <div className="text-zinc-400 text-xs">
                                        {label}
                                    </div>
                                    <div className="text-lg font-semibold">
                                        {value}
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </aside>
            </div>
        </div>
    );
}

export default Build;
