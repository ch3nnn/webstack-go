/**
 * 将扁平化列表转换为树形结构
 * @param {Array} list - 扁平化列表
 * @param {number} pid - 父节点 ID，默认为 0（根节点）
 * @returns {Array} - 树形结构
 */
function buildTree(list, pid = 0) {
    return list
        .filter(item => item.pid === pid) // 过滤出当前层级的节点
        .map(item => ({
            ...item,
            children: buildTree(list, item.id) // 递归查找子节点
        }));
}
