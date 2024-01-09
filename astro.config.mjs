import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import react from '@astrojs/react';
import sitemap from '@astrojs/sitemap';
import remarkToc from 'remark-toc';
import rehypeSlug from 'rehype-slug';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';
// https://astro.build/config
import astroRemark from "@astrojs/markdown-remark";
export default defineConfig({
	site: 'https://example.com',
	integrations: [mdx(), sitemap(), react()],
	
	markdownOptions: {
        render: [
            astroRemark,
            {
                rehypePlugins: [
                    "rehype-slug",
                    [
                        "rehype-autolink-headings",
                        { behavior: "append"},
                    ],
                    [
                        "rehype-toc",
                        { headings: ["h1", "h2"] }
                    ]
                ],
            },
        ],
    },
});
