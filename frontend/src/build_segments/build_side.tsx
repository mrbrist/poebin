import { type buildData } from "../api/getBuild";
import { timeAgo } from "../helpers/timeAgo";

function BuildSide({ build }: { build: buildData | undefined }) {
    const buildClass =
        build?.Data.Build.AscendClassName == "None"
            ? build.Data.Build.ClassName
            : build?.Data.Build.AscendClassName;

    function copyBuildToClipboard() {
        if (build?.Raw) {
            navigator.clipboard.writeText(build.Raw);
        }
    }
    return (
        <aside className="col-span-3 space-y-6">
            <div className="rounded-2xl bg-slate-900 border border-slate-400 p-4 shadow-lg">
                <h2 className="text-lg font-semibold mb-4">Build Info</h2>
                {buildClass ? (
                    <div className="w-full flex justify-center">
                        <img
                            // This may need to change to local at some point if rate limits happen
                            src={`https://assets.poe.ninja/poe1/classes/${buildClass.toLowerCase()}.webp`}
                            alt={buildClass}
                            className="w-auto rounded-xl border border-slate-700 mb-4"
                        />
                    </div>
                ) : null}

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
                    <div>
                        <button
                            className="inline-flex items-center gap-2 p-2 mt-4 mb-2 cursor-pointer bg-slate-700 hover:bg-slate-600 border border-slate-400 rounded-sm"
                            onClick={copyBuildToClipboard}
                        >
                            <span id="default">
                                <svg
                                    className="w-4 h-4 text-fg-brand"
                                    aria-hidden="true"
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                >
                                    <path
                                        stroke="currentColor"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth="2"
                                        d="M15 4h3a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3m0 3h6m-6 7 2 2 4-4m-5-9v4h4V3h-4Z"
                                    />
                                </svg>
                            </span>
                            <span id="success" className=""></span>
                            <span>Copy PoB Code</span>
                        </button>
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
