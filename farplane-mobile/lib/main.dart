import 'package:flutter/material.dart';
import 'purchase_transaction.dart';

void main() {
  runApp(const FarplaneApp());
}

class FarplaneApp extends StatelessWidget {
  const FarplaneApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    final ThemeData theme = ThemeData();
    return MaterialApp(
      title: 'Farplane',
      theme: theme.copyWith(
        colorScheme: theme.colorScheme.copyWith(
          primary: Colors.grey,
          secondary: Colors.black,
        ),
      ),
      home: const MyHomePage(title: 'Farplane'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: SafeArea(
        child: ListView.builder(
          itemCount: 1,
          itemBuilder: (BuildContext context, int index) {
            return buildPurchaseTransactionCard(
                PurchaseTransaction.static_test_transactions[0]);
          },
        ),
      ),
    );
  }

  Widget buildPurchaseTransactionCard(PurchaseTransaction purchaseTransaction) {
    return Card(
      child: Column(
        children: <Widget>[
          Image(image: AssetImage(purchaseTransaction.imageUrl)),
          Text(purchaseTransaction.description),
        ],
      ),
    );
  }
}
