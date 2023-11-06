import { NextApiRequest, NextApiResponse } from "next";

export default function Up(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method == "HEAD") {
        res.status(200).end();
        return;
    }
    if (req.method != "POST") {
        res.status(405).json({ message: "use POST or HEAD instead." });
        return;
    }

    // TODO generate key and extract filename from request

    res.status(200).json({ key: "coolkey" });
}
