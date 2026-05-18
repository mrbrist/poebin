import { useEffect, useState } from "react";
import type { buildData } from "../api/getBuild";
import rawTreeData from "../../tree_data/328.json";

function BuildTree({ build }: { build: buildData | undefined }) {
    const trees = build?.Data?.Tree?.Specs ?? [];

    const [selectedTreeTitle, setSelectedTreeTitle] = useState<string>("");

    // initialize once build loads
    useEffect(() => {
        if (trees.length > 0 && !selectedTreeTitle) {
            setSelectedTreeTitle(trees[0].Title);
        }
    }, [trees]);

    if (!build || trees.length === 0) return <div>Loading...</div>;

    const selectedTree = trees.find((s: any) => s.Title === selectedTreeTitle);

    if (!selectedTree) return null;

    console.log(rawTreeData);

    function lookupNode(node: string) {
        const data = rawTreeData.nodes?.[node];

        // Always return a safe object
        if (!data) {
            return {
                name: "Unknown Node",
                text: "",
                ascendancyNode: false,
                isNotable: false,
            };
        }

        return {
            name: data.name ?? "Unnamed Node",
            text: Array.isArray(data.stats) ? data.stats.join(", ") : "",
            ascendancyNode: !!data.ascendancyName,
            isNotable: !!data.isNotable,
        };
    }

    function getPassiveSkillData(tree: { Sockets?: any[]; Nodes?: string }) {
        // Prevent crashes from malformed tree data
        if (!tree?.Nodes) {
            return {
                treeNodes: [],
                ascendancyNodes: [],
                socketCount: 0,
            };
        }

        const nodes = tree.Nodes.split(",");
        const socketCount = Array.isArray(tree.Sockets)
            ? tree.Sockets.length
            : 0;

        const treeNodes = [];
        const ascendancyNodes = [];

        for (let i = 0; i < nodes.length; i++) {
            const n = lookupNode(nodes[i]);

            if (!n?.isNotable) continue;

            if (n.ascendancyNode) {
                ascendancyNodes.push(n);
            } else {
                treeNodes.push(n);
            }
        }

        return {
            treeNodes,
            ascendancyNodes,
            socketCount,
        };
    }

    let treeData = getPassiveSkillData(selectedTree);

    return (
        <div className="block text-left">
            {trees.length > 1 && (
                <select
                    value={selectedTreeTitle}
                    onChange={(e) => setSelectedTreeTitle(e.target.value)}
                    className="mb-4 rounded border p-2"
                >
                    {trees.map((set: any) => (
                        <option key={set.Title} value={set.Title}>
                            {set.Title}
                        </option>
                    ))}
                </select>
            )}

            <div className="rounded-2xl bg-slate-900 border border-slate-400 p-6 shadow-lg mt-6">
                <div className="mb-3 text-sm text-zinc-400">
                    This tree has{" "}
                    <span className="text-fuchsia-400 font-semibold">
                        {treeData.socketCount}
                    </span>{" "}
                    jewel sockets allocated.
                </div>

                {/* Ascendancy */}
                <div className="mb-6">
                    <h3 className="text-md font-semibold text-zinc-200 mb-3">
                        Ascendancy
                    </h3>

                    {treeData.ascendancyNodes.length > 0 ? (
                        <div className="space-y-3">
                            {treeData.ascendancyNodes.map((node, index) => (
                                <div
                                    key={index}
                                    className="bg-slate-800 p-4 rounded-xl border border-amber-700"
                                >
                                    <div className="text-amber-400 font-medium mb-1">
                                        {node.name}
                                    </div>
                                    <div className="text-sm text-zinc-400">
                                        {node.text}
                                    </div>
                                </div>
                            ))}
                        </div>
                    ) : (
                        <div className="text-sm text-zinc-500">
                            No ascendancy passives allocated.
                        </div>
                    )}
                </div>

                {/* Regular */}
                <div>
                    <h3 className="text-md font-semibold text-zinc-200 mb-3">
                        Notables
                    </h3>

                    {treeData.treeNodes.length > 0 ? (
                        <div className="space-y-3">
                            {treeData.treeNodes.map((node, index) => (
                                <div
                                    key={index}
                                    className="bg-slate-800 p-4 rounded-xl border border-slate-700"
                                >
                                    <div className="text-fuchsia-400 font-medium mb-1">
                                        {node.name}
                                    </div>
                                    <div className="text-sm text-zinc-400">
                                        {node.text}
                                    </div>
                                </div>
                            ))}
                        </div>
                    ) : (
                        <div className="text-sm text-zinc-500">
                            No notable passives allocated.
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

export default BuildTree;
