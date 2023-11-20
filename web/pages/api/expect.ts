import { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse,
) {
    // Generate keyword
    // Store keyword -> IP in KV store
    // Send keyword to user

    res.status(200).json({ keyword: "cat", ok: true });
}
