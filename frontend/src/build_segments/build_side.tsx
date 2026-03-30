import { type buildData } from "../api/getBuild";
import { timeAgo } from "../helpers/timeAgo";

function BuildSide({ build }: { build: buildData | undefined }) {
    return (
        <aside className="col-span-3 space-y-6">
            <div className="rounded-2xl bg-slate-900 border border-slate-400 p-4 shadow-lg">
                <h2 className="text-lg font-semibold mb-4">Build Info</h2>
                <div className="space-y-2 text-sm text-zinc-400">
                    <div>
                        <span className="text-zinc-200">Class:</span>{" "}
                        {build?.Data.Build.ClassName}
                    </div>
                    <div>
                        <span className="text-zinc-200">Ascendancy:</span>{" "}
                        {build?.Data.Build.AscendClassName}
                    </div>
                    <div>
                        <span className="text-zinc-200">Level: </span>
                        {build?.Data.Build.Level}
                    </div>
                    <div>{build ? timeAgo(build?.LastModified) : null}</div>
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
                            <div className="text-zinc-400 text-xs">{label}</div>
                            <div className="text-lg font-semibold">{value}</div>
                        </div>
                    ))}
                </div>
            </div>
        </aside>
    );
}

export default BuildSide;
