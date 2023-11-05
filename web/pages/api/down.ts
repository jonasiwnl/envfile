import { NextApiRequest, NextApiResponse } from "next";

export default function Down(
    req: NextApiRequest,
    res: NextApiResponse
) {
    if (req.method == "HEAD") {
        res.status(200).end();
        return;
    }
    if (req.method != "GET") {
        res.status(405).json({ message: "use GET or HEAD instead." });
        return;
    }
    const key = req.query.key;
    if (typeof key != "string") { // TODO array of strings?
        res.status(400).json({ message: "key must be defined." });
        return;
    }

    // TODO lookup and respond with IP
    res.status(200).json({ key, ip: "idk, sorry" });
}
