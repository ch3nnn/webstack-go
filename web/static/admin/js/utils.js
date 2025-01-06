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


/**
 * 递归渲染树形结构到 <select>
 * @param {Array} data - 树形结构数据
 * @param {HTMLElement} parent - 父元素
 * @param {number} level - 当前层级（用于缩进）
 * @param {Array} selectedIds - 需要选中的选项 ID 列表
 */
function renderTreeToSelect(data, parent, level = 0, selectedIds = []) {
    data.forEach(item => {
        if (item.children && item.children.length > 0) {
            // 如果是分组，创建 <optgroup>
            const optgroup = document.createElement('optgroup');
            optgroup.label = item.name;
            parent.appendChild(optgroup);

            // 递归渲染子节点
            renderTreeToSelect(item.children, optgroup, level + 1, selectedIds);
        } else {
            // 如果是选项，创建 <option>
            const option = document.createElement('option');
            option.value = item.id;
            option.text = ' '.repeat(level * 4) + item.name; // 缩进显示

            // 如果当前选项 ID 在 selectedIds 中，则设置为选中
            if (selectedIds.includes(item.id)) {
                option.selected = true;
            }

            parent.appendChild(option);
        }
    });
}
