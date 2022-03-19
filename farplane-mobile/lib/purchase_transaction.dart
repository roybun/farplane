class PurchaseTransaction {
  String description;
  String imageUrl;

  PurchaseTransaction(
    this.description,
    this.imageUrl,
  );

  static List<PurchaseTransaction> static_test_transactions = [
    PurchaseTransaction('Fuck You Veronica Bitch', 'assets/fdathoe.jpg'),
  ];
}
