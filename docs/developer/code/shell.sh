/// [merge-branch]
#!/bin/bash
function red() {
    echo -e "\033[31m$1\033[0m"
}

function yellow() {
    echo -e "\033[33m$1\033[0m"
}

function green() {
    echo -e "\033[32m$1\033[0m"
}

CURRENT_BRANCH=$(git symbolic-ref -q --short HEAD);
if [[ $CURRENT_BRANCH =~ "test" ]]; then 
    red "无法将`test`合并到其他分支"
    exit 1
fi

if [[ $CURRENT_BRANCH =~ "preline" ]]; then 
    red "无法将`preline`合并到其他分支"
    exit 1
fi

green "Current branch: $CURRENT_BRANCH"

if [ -n "$1" ]; then
    MERGE_TO_BRANCH=$1
else
    red "Please specify merge to branch"
    exit 1
fi

if [ "$CURRENT_BRANCH"x = "$MERGE_TO_BRANCH"x ]; then
    red "can't  merge to self"
    exit 1
fi

CURRENT_BRANCH=$(git symbolic-ref -q --short HEAD);
green "Merge: $CURRENT_BRANCH -> $MERGE_TO_BRANCH";
echo ""

yellow "1、切换$MERGE_TO_BRANCH"
git checkout $MERGE_TO_BRANCH
echo ""

yellow "2、Pull $MERGE_TO_BRANCH 代码"
git pull origin $MERGE_TO_BRANCH
echo ""

yellow "3、Pull $CURRENT_BRANCH 代码"
git pull origin $CURRENT_BRANCH
echo ""

yellow "4、Push To $MERGE_TO_BRANCH"
git push origin $MERGE_TO_BRANCH
echo ""

yellow "切换回 $CURRENT_BRANCH"
git checkout $CURRENT_BRANCH
echo ""

green "Successfully"
echo ""
/// [merge-branch]

