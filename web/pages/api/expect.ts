import { NextApiRequest, NextApiResponse } from "next";
import { createClient } from "redis";

const redis = createClient({
    url: process.env.REDIS,
});

const keywords = [
    "cat",
    "dog",
    "apple",
    "banana",
    "orange",
    "pear",
    "bird",
    "fish",
    "turtle",
    "snake",
    "lizard",
    "frog",
    "pear",
];

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse,
) {
    if (req.method !== "GET") return res.status(405).end();

    await redis.connect();

    // TODO check if keyword is already in redis
    const keyword = keywords[Math.floor(Math.random() * keywords.length)];
    const ok = true;

    redis.set(keyword, req.socket.remoteAddress || "");

    void redis.disconnect();
    res.status(200).json({ keyword, ok });
}
