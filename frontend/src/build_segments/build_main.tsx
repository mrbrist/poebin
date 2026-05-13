import { type buildData } from "../api/getBuild";
import { useEffect, useState } from "react";
import BuildItems from "./items";
import BuildTree from "./tree";
import BuildSkills from "./skills";
import { useNavigate, useParams } from "react-router-dom";

function BuildMain({ build }: { build: buildData | undefined }) {
    const navigate = useNavigate();
    const { id, tab } = useParams();

    const getValidTab = (
        currentTab: string | undefined,
    ): "items" | "tree" | "skills" => {
        if (
            currentTab === "items" ||
            currentTab === "tree" ||
            currentTab === "skills"
        ) {
            return currentTab;
        }

        return "items";
    };

    const [activeTab, setActiveTab] = useState<"items" | "tree" | "skills">(
        getValidTab(tab),
    );

    useEffect(() => {
        setActiveTab(getValidTab(tab));
    }, [tab]);

    const handleTabChange = (newTab: "items" | "tree" | "skills") => {
        setActiveTab(newTab);
        navigate(`/${id}/${newTab}`);
    };

    return (
        <div className="col-span-9 rounded-2xl border border-slate-400 p-4">
            <div className="mb-4 flex gap-4 border-b border-slate-300 pb-2">
                <button
                    className={`rounded-t-lg px-3 py-1 ${
                        activeTab === "items"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => handleTabChange("items")}
                >
                    Items
                </button>

                <button
                    className={`rounded-t-lg px-3 py-1 ${
                        activeTab === "tree"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => handleTabChange("tree")}
                >
                    Tree
                </button>

                {/* Uncomment when skills tab is ready */}
                {/* 
                <button
                    className={`rounded-t-lg px-3 py-1 ${
                        activeTab === "skills"
                            ? "bg-slate-700 font-semibold"
                            : "text-zinc-400 hover:bg-slate-800"
                    }`}
                    onClick={() => handleTabChange("skills")}
                >
                    Skills
                </button>
                */}
            </div>

            {activeTab === "items" && <BuildItems build={build} />}
            {activeTab === "tree" && <BuildTree build={build} />}
            {activeTab === "skills" && <BuildSkills build={build} />}
        </div>
    );
}

export default BuildMain;
