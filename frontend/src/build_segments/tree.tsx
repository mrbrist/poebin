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

        // Prevent crashes if node doesn't exist
        if (!data) {
            return {
                name: "Unknown Node",
                text: "",
                ascendancyNode: false,
            };
        }

        return {
            name: data.name ?? "Unnamed Node",
            text: data.stats?.join(", ") ?? "",
            ascendancyNode: data.ascendancyName ? true : false,
            isNotable: data.isNotable ? true : false,
        };
    }

    function getPassiveSkillData(tree: { Sockets: any; Nodes: string }) {
        const nodes = tree.Nodes.split(",");
        const socketCount = tree.Sockets.length;

        const treeNodes = [];
        const ascendancyNodes = [];

        for (let i = 0; i < nodes.length; i++) {
            const n = lookupNode(nodes[i]);

            // Only include notable nodes
            if (!n.isNotable) continue;

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

            <div>
                This tree has{" "}
                <span className="text-fuchsia-500">{treeData.socketCount}</span>{" "}
                sockets allocated.
                {treeData.treeNodes.map((node) => (
                    <div>
                        {node.name}: {node.text}
                    </div>
                ))}
                ----
                {treeData.ascendancyNodes.map((node) => (
                    <div>
                        {node.name}: {node.text}
                    </div>
                ))}
            </div>
        </div>
    );
}

export default BuildTree;
