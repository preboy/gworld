
// ����һ�����ڵ���0��С��1��α�����
exports.getRandom = () => {
  return Math.random();
}


// ����һ������min��max֮��������
exports.getRandomArbitrary = (min, max) => {
  return Math.random() * (max - min) + min;
}

// ����һ������min��max֮������������
// Using Math.round() will give you a non-uniform distribution!
exports.getRandomInt = (min, max) => {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
