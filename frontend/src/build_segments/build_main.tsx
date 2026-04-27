import { type buildData } from "../api/getBuild";
import { useState } from "react";
import BuildItems from "./items";
import BuildTree from "./tree";
import BuildSkills from "./skills";

function BuildMain({ build }: { build: buildData | undefined }) {
    const [activeTab, setActiveTab] = useState<"items" | "tree" | "skills">(
        "items",
    );

    return (
        <div className="col-span-9 rounded-2xl border border-slate-400 p-4">
            <div className="flex gap-4 mb-4 border-b border-slate-300 pb-2">
                <button
                    className={`px-3 py-1 rounded-t-lg ${
                        activeTab === "items"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => setActiveTab("items")}
                >
                    Items
                </button>

                <button
                    className={`px-3 py-1 rounded-t-lg ${
                        activeTab === "tree"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => setActiveTab("tree")}
                >
                    Tree
                </button>

                {/* <button
                    className={`px-3 py-1 rounded-t-lg ${
                        activeTab === "skills"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => setActiveTab("skills")}
                >
                    Skills
                </button> */}
            </div>
            {activeTab === "items" && <BuildItems build={build} />}
            {activeTab === "tree" && <BuildTree build={build} />}
            {activeTab === "skills" && <BuildSkills build={build} />}
        </div>
    );
}

export default BuildMain;
