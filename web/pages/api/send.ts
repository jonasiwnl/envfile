import { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse,
) {
    // Lookup keyword -> IP in KV store
    // Send IP to user
    res.status(200).json({ ip: "0.0.0.0:8080", ok: true });
}
