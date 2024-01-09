import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import react from '@astrojs/react';
import sitemap from '@astrojs/sitemap';
import remarkToc from 'remark-toc';
import rehypeSlug from 'rehype-slug';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';

import astroRemark from '@astrojs/markdown-remark';
export default defineConfig({
	site: 'https://example.com',
	integrations: [mdx(), sitemap()],
	markdown: {

		// 应用于 .md 和 .mdx 文件
		remarkPlugins: [[remarkToc, { heading: "contents", headings: ['h1', 'h2'] }]],
		rehypePlugins: [rehypeSlug, [rehypeAutolinkHeadings, { behavior: 'append' }]],
		gfm: true,
		shikiConfig: {
			// 选择 Shiki 内置的主题（或添加你自己的主题）
			// https://github.com/shikijs/shiki/blob/main/docs/themes.md
			theme: 'dracula',
		},
	}
});
