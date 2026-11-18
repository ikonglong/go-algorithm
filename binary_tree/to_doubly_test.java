package binary_tree

// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
// 把二叉树转换为双向链表，顺序为中序，就地转换。left 指向前驱，right 指向后继

class Solution {
    public Node treeToDoublyList(Node root) {
        if (root == null) {
            return null;
        }

        Node dummyHead = new Node(0);
        Node[] tail = {dummyHead};

        inorder(root, tail);

        Node head = dummyHead.right;
        head.left = tail[0];
        tail[0].right = head;

        return head;
    }

    private void inorder(Node node, Node[] tail) {
        if (node == null) {
            return;
        }

        inorder(node.left, tail);

        // right 作为 next，left 作为 prev
        tail[0].right = node;
        node.left = tail[0];
        tail[0] = node;

        inorder(node.right, tail);
    }
}
