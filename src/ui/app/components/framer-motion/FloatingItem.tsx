"use client";

import { motion } from "framer-motion";

interface FloatingItemProps {
    children: React.ReactNode;
    className?: string;  // Optional className prop added here
}

const FloatingItem = ({ children, className }: FloatingItemProps) => {

    const arr1 = Array.from({length: 6}, () => Math.random() * (4 - -4) + -4);
    const arr2 = Array.from({length: 6}, () => Math.random() * (4 - -4) + -4);
    const arr3 = Array.from({length: 6}, () => Math.random() * (4 - -4) + -4);


    return (
        <motion.div
            className={className}
            initial={{ x: 0, y: 0 }}
            animate={{
                x: arr1,  // Randomized horizontal movement
                y: arr2,   // Randomized vertical movement
                rotate: arr3, // Subtle rotation for a natural feel
            }}
            transition={{
                duration: 4,              // Longer duration for slower movement
                ease: "easeInOut",        // Smooth easing curve
                repeat: Infinity,         // Infinite loop
                repeatType: "mirror",     // Ping-pong effect for back-and-forth movement
            }}
            style={{ display: "inline-block" }}
        >
            {children}
        </motion.div>
    );
};

export default FloatingItem;