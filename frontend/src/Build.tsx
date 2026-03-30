import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { getBuild, type buildData } from "./api/getBuild";

// Build Segments
import BuildSide from "./build_segments/build_side";
import BuildMain from "./build_segments/build_main";

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
                <BuildSide build={build} />
                <BuildMain build={build} />
            </div>
        </div>
    );
}

export default Build;
