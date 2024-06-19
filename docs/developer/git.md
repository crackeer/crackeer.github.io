## 删除git中大文件+记录
- https://blog.csdn.net/HappyRocking/article/details/89313501

```bash
git rev-list --all | xargs -rL1 git ls-tree -r --long | sort -uk3 | sort -rnk4 | head -10
git filter-branch --tree-filter "rm -f {filepath}" -- --all
git push -f --all
```